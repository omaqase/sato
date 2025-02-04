import dotenv from "dotenv";
import express from "express";
import bodyParser from "body-parser";
import OpenAI from "openai";
import cors from "cors";
import { sleep } from "bun";

dotenv.config();

const openai = new OpenAI({
  baseURL: "https://openrouter.ai/api/v1",
  apiKey:
    "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", // Replace with your actual API key
});

async function retryWithBackoff(fn, retries = 5, delay = 1000) {
  for (let attempt = 0; attempt < retries; attempt++) {
    try {
      return await fn();
    } catch (error) {
      if (error.code === 429 && attempt < retries - 1) {
        console.log(`Rate limited. Retrying in ${delay}ms...`);
        await new Promise((resolve) => setTimeout(resolve, delay));
        delay *= 2;
      } else {
        throw error;
      }
    }
  }
}

const app = express();
const port = process.env.PORT || 3000;

app.use(bodyParser.json());
app.use(cors());

app.post("/ask", async (req, res) => {
  const { prompt } = req.body;

  if (!prompt) {
    return res.status(400).json({ error: "Prompt is required!" });
  }

  res.setHeader("Content-Type", "text/plain");
  res.setHeader("Transfer-Encoding", "chunked");

  try {
    const completion = await retryWithBackoff(() =>
      openai.chat.completions.create({
        model: "sophosympatheia/rogue-rose-103b-v0.2:free",
        stream: true,
        messages: [
          {
            role: "user",
            content: [
              {
                type: "text",
                text: prompt,
              },
            ],
          },
        ],
      })
    );

    try {
      for await (const chunk of completion) {
        sleep(1000);
        const choice = chunk.choices[0];
        if (choice && choice.delta && choice.delta.content) {
          const token = choice.delta.content;
          console.log(token);
          res.write(token);
        }

        if (choice.finish_reason === "stop") {
          console.log("Stream finished.");
          break;
        }
      }
    } catch (err) {
      console.error("Error processing stream:", err);
      res.status(500).end("An error occurred while streaming the response.");
    }

    res.end();
  } catch (error) {
    console.error("Error during AI response generation:", error);
    res.status(500).end("An error occurred while generating the response.");
  }
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});

require("dotenv").config();
const express = require("express");
const bodyParser = require("body-parser");
const OpenAI = require("openai");

const app = express();
const port = process.env.PORT || 3000;

// Middleware to parse JSON
app.use(bodyParser.json());

// Initialize OpenAI client
const openai = new OpenAI({
  baseURL: "https://openrouter.ai/api/v1",
  apiKey: process.env.OPENROUTER_API_KEY,
  defaultHeaders: {
    "HTTP-Referer": process.env.YOUR_SITE_URL, // Optional. Site URL for rankings on openrouter.ai.
    "X-Title": process.env.YOUR_SITE_NAME, // Optional. Site title for rankings on openrouter.ai.
  },
});

// Helper function to check if the query is within the allowed topics
function isValidQuery(userMessage) {
  const validKeywords = [
    "заказ",
    "доставка",
    "статус",
    "отмена",
    "профиль",
    "ошибка",
    "не работает",
    "проблема",
  ];
  return validKeywords.some((keyword) =>
    userMessage.toLowerCase().includes(keyword)
  );
}

// Route to handle user queries
app.post("/ask", async (req, res) => {
  const { message } = req.body;

  if (!isValidQuery(message)) {
    return res.json({
      response:
        "Извините, я могу помочь только с вопросами о заказах, профиле, информации о маркетплейсе Satori или ошибках на сайте. Пожалуйста, задайте вопрос по этим темам! 😊",
    });
  }

  try {
    const completion = await openai.chat.completions.create({
      model: "google/gemini-2.0-flash-exp:free",
      messages: [
        {
          role: "system",
          content: `
            Ты - Алуа, помощник для Satori Marketplace. Твоя задача - отвечать только на вопросы, связанные с заказами, профилем пользователя, информацией о маркетплейсе, пунктами самовывоза и ошибками на сайте.
            Если вопрос не относится к этим темам, ответь: "Извините, я могу помочь только с вопросами о заказах, профиле, информации о маркетплейсе Satori или ошибках на сайте. Пожалуйста, задайте вопрос по этим темам! 😊"
          `,
        },
        {
          role: "user",
          content: message,
        },
      ],
    });

    res.json({ response: completion.choices[0].message });
  } catch (error) {
    console.error("Error communicating with OpenAI:", error);
    res.status(500).json({
      error: "Произошла ошибка при обработке запроса. Попробуйте позже.",
    });
  }
});

// Start the server
app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});

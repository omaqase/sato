import openai from "./openaiClient";
import { generateIntentPrompt, generateResponsePrompt } from "./promptEngineer";

export async function handleUserQuery(
  question: string,
  context?: any
): Promise<string> {
  try {
    // Step 1: Analyze intent using OpenAI
    const intentPrompt = generateIntentPrompt(question);
    openai.beta.assistants;
    const intentCompletion = await openai.chat.completions.create({
      model: "google/gemini-2.0-flash-exp:free",
      messages: [{ role: "user", content: intentPrompt }],
      max_tokens: 100, // Ограничиваем размер ответа
    });

    if (!intentCompletion.choices || intentCompletion.choices.length === 0) {
      return "Извините, не удалось проанализировать ваш запрос. Попробуйте позже!";
    }

    let intentData;
    try {
      // Extract and parse JSON from the raw response
      intentData = extractJson(
        intentCompletion.choices[0].message.content || ""
      );
    } catch (parseError) {
      console.error("Error parsing JSON:", parseError);
      return "Извините, произошла ошибка при анализе вашего запроса. Пожалуйста, попробуйте позже.";
    }

    const { intent, orderId } = intentData;

    // Validate order ID match
    if (
      context?.type === "order" &&
      orderId &&
      orderId !== context.data.order_id
    ) {
      return `Извините, но заказ ${orderId} не найден. Пожалуйста, проверьте номер заказа. 🤔`;
    }

    // Step 2: Generate response based on intent
    const responsePrompt = generateResponsePrompt(intent, context);

    const responseCompletion = await openai.chat.completions.create({
      model: "google/gemini-2.0-flash-exp:free",
      messages: [{ role: "user", content: responsePrompt }],
      max_tokens: 100, // Ограничиваем размер ответа
    });

    if (
      !responseCompletion.choices ||
      responseCompletion.choices.length === 0
    ) {
      return "Извините, не удалось получить ответ. Попробуйте позже!";
    }

    return (
      responseCompletion.choices[0].message.content ||
      "Нет данных для отображения."
    );
  } catch (error) {
    console.error("Error while processing the request:", error);

    if (error instanceof Error) {
      console.error("Error details:", error.message);
    }

    return "Произошла ошибка при обработке вашего запроса. Пожалуйста, попробуйте позже.";
  }
}

// Helper function to extract JSON from raw text
function extractJson(rawText: string): any {
  try {
    // Ищем JSON в тексте с помощью регулярного выражения
    const jsonMatch = rawText.match(/```json\s*({[\s\S]*?})\s*```/);
    if (!jsonMatch) {
      throw new Error("JSON block not found in the response");
    }

    // Парсим JSON
    return JSON.parse(jsonMatch[1]);
  } catch (error) {
    console.error("Error extracting JSON:", error);
    console.error("Raw OpenAI response:", rawText);
    throw error;
  }
}

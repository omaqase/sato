import { Bot, type Context } from "https://deno.land/x/grammy/mod.ts";
import { serve } from "https://deno.land/std/http/server.ts";
import { NotificationGrpcServer } from "./grpc_server.ts";

interface Receipt {
    orderId: string;
    items: {
        name: string;
        price: number;
        quantity: number;
    }[];
    total: number;
    customerName: string;
    date: string;
}

interface Notification {
    type: "order_status" | "payment_status" | "delivery_status";
    message: string;
    orderId: string;
}

class TelegramNotificationService {
    private bot: Bot;
    private chatIds: Set<number> = new Set();

    constructor(token: string) {
        this.bot = new Bot(token);
        this.setupBot();
    }

    private setupBot() {
        this.bot.command("start", (ctx) => {
            this.chatIds.add(ctx.chat.id);
            ctx.reply("Бот успешно подключен! Вы будете получать уведомления о заказах.");
        });

        this.bot.command("id", (ctx) => {
            ctx.reply(`Ваш Chat ID: ${ctx.chat.id}`);
        });
    }

    async sendReceipt(chatId: number, receipt: Receipt) {
        const message = this.formatReceipt(receipt);
        await this.bot.api.sendMessage(chatId, message, { parse_mode: "HTML" });
    }

    async sendNotification(chatId: number, notification: Notification) {
        const message = this.formatNotification(notification);
        await this.bot.api.sendMessage(chatId, message, { parse_mode: "HTML" });
    }

    private formatReceipt(receipt: Receipt): string {
        let message = `🧾 <b>Чек #${receipt.orderId}</b>\n`;
        message += `👤 Покупатель: ${receipt.customerName}\n`;
        message += `📅 Дата: ${receipt.date}\n\n`;
        message += `📋 Товары:\n`;
        
        receipt.items.forEach(item => {
            message += `- ${item.name} x${item.quantity} = ${item.price * item.quantity}₽\n`;
        });
        
        message += `\n💰 Итого: ${receipt.total}₽`;
        return message;
    }

    private formatNotification(notification: Notification): string {
        return `🔔 <b>Уведомление</b>\n📦 Заказ #${notification.orderId}\n${notification.message}`;
    }

    startBot() {
        this.bot.start();
    }
}

// Создаем экземпляр сервиса
const notificationService = new TelegramNotificationService("7452367801:AAHq3HmhCNI7zttlMUhnv24UShtL0NGK1Xo");

// В конце файла заменим старый код запуска на:
const startServers = async () => {
    // Запускаем HTTP сервер
    serve({
        port: 3000,
        handler: async (request: Request) => {
            if (request.method !== "POST") {
                return new Response("Method not allowed", { status: 405 });
            }

            try {
                const body = await request.json();
                const chatId = body.chatId;

                if (!chatId) {
                    return new Response("Chat ID is required", { status: 400 });
                }

                if (request.url.endsWith("/send-receipt")) {
                    await notificationService.sendReceipt(chatId, body.receipt);
                } else if (request.url.endsWith("/send-notification")) {
                    await notificationService.sendNotification(chatId, body.notification);
                } else {
                    return new Response("Invalid endpoint", { status: 404 });
                }

                return new Response("Message sent successfully", { status: 200 });
            } catch (error) {
                return new Response(`Error: ${error.message}`, { status: 500 });
            }
        },
    });
    console.log("HTTP server started on port 3000");

    // Запускаем gRPC сервер
    const grpcServer = new NotificationGrpcServer(notificationService);
    await grpcServer.start(50051);

    // Запускаем бота
    notificationService.startBot();
    console.log("Telegram bot started");
};

startServers();
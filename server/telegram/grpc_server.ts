import { GrpcServer } from "https://deno.land/x/grpc_basic@0.4.7/server.ts";
import { TelegramNotificationService } from "./main.ts";

export class NotificationGrpcServer {
    private server: GrpcServer;
    private notificationService: TelegramNotificationService;
    private listener: Deno.Listener | null = null;

    constructor(notificationService: TelegramNotificationService) {
        this.server = new GrpcServer();
        this.notificationService = notificationService;
        this.setupServices();
    }

    private setupServices() {
        const protoPath = new URL("./notification.proto", import.meta.url);
        this.server.addService(protoPath, {
            SendReceipt: async (req: any) => {
                try {
                    await this.notificationService.sendReceipt(req.chat_id, req.receipt);
                    return { success: true, message: "Receipt sent successfully" };
                } catch (error) {
                    return { success: false, message: error.message };
                }
            },
            SendNotification: async (req: any) => {
                try {
                    await this.notificationService.sendNotification(req.chat_id, req.notification);
                    return { success: true, message: "Notification sent successfully" };
                } catch (error) {
                    return { success: false, message: error.message };
                }
            }
        });
    }

    async start(port: number) {
        const addr = `0.0.0.0:${port}`;
        this.listener = Deno.listen({ port, hostname: "0.0.0.0" });
        
        console.log(`gRPC server started on ${addr}`);
        
        for await (const conn of this.listener) {
            this.server.handle(conn);
        }
    }

    async stop() {
        if (this.listener) {
            this.listener.close();
            this.listener = null;
        }
    }
}
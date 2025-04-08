const { Kafka } = require("kafkajs");
const cron = require('node-cron')

const kafka = new Kafka({
    clientId: 'my-app',
    brokers: ['localhost:29092'],
});
const topic = 'sample';

// 구독
const consumer = kafka.consumer({groupId: 'test-group'})

async function consumeMessage() {
    try {
        await consumer.connect();

        await consumer.subscribe({topic, fromBeginning: true})

        await consumer.run({
            eachMessage: async ({topic, partition, message}) => {
                console.log({
                    partition,
                    offset: message.offset,
                    value: message.value.toString(),
                });
            },
        });
    } catch (error) {
        console.error("Error consuming message:", error)
    }
}

consumeMessage();

// 발행
/*const producer = kafka.producer();

cron.schedule("*!/5 * * * * *", async () => {
    try {
        await producer.connect();

        await producer.send({
            topic,
            messages: [{value: "Message from cron job!"}],
        });
        console.log("Message sent from cron job!");
    } catch (error) {
        console.error("Error sending message from cron job:", error);
    }
})*/


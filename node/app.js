const { Kafka } = require("kafkajs");

const kafka = new Kafka({
  brokers: ["localhost:9092"], // Use port 9092 for Kafka
  clientId: "example-consumer",
});

const consumer = kafka.consumer({ groupId: "test-group" });
const admin = kafka.admin();

const run = async () => {
  // Producing

  // Consuming
  await consumer.connect();
  const topics = await admin.listTopics();
  console.log(topics);
  await consumer.subscribe({
    topic: "fullfillment.your_database_server_name.db0.table_name",
    fromBeginning: true,
  });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      console.log({
        partition,
        offset: message.offset,
        value: message.value.toString(),
      });
    },
  });
};

run().catch(console.error);

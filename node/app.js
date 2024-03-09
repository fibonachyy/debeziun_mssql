const { Kafka } = require("kafkajs");

const kafka = new Kafka({
  brokers: ["localhost:9092"], // Use port 9092 for Kafka
  clientId: "example-consumer",
});

const consumer = kafka.consumer({ groupId: "test-group" });
const admin = kafka.admin();

const run = async () => {
  // Connect the consumer
  await consumer.connect();

  // List of topics to subscribe to
  await admin.connect();

  const topics = await admin.listTopics();

  console.log("List of topics:", topics);
  // Subscribe to each topic in the list
  await Promise.all(
    topics.map(async (topic) => {
      await consumer.subscribe({ topic: "sqlserver.dbo.table_name" });
      console.log(`Subscribed to topic: ${topic}`);
    })
  );

  // Start consuming messages from all subscribed topics
  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      console.log({
        topic,
        partition,
        offset: message.offset,
        value: message.value.toString(),
      });
    },
  });
};

run().catch(console.error);

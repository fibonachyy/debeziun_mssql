const { Kafka } = require("kafkajs");

const kafka = new Kafka({
  brokers: ["localhost:9092"], // Use port 9092 for Kafka
  clientId: "example-consumer",
});

const consumer = kafka.consumer({ groupId: "test-group" });

const run = async () => {
  // Connect the consumer
  await consumer.connect();

  // List of topics to subscribe to
  const topics = [
    "sqlserver.dbo.table_name",
    // Add more topics as needed
  ];

  // Subscribe to each topic in the list
  await Promise.all(
    topics.map(async (topic) => {
      await consumer.subscribe({ topic });
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

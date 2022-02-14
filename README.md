# rabbitmq_task

## this is a task aimed at practicing RabbitMQ within the framework of my iTechArt traineeship.

### the program has 3 layers:
- Collectors collects links from a defined webpage and sends them to the Worker via RabbitMQ queue;
- Worker collects the URL of each link received from a queue, gets the title of that URL and sends both the URL and the Title to the Saver via 2 different RabbitMQ routing keys;
- Saver gets URL and Title from 2 different routing keys and saves them into a .txt file.

### the latest version of the program features direct exchange.

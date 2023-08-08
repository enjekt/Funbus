# Hexabus

The hexabus is designed for use in hexagonal microservices architectures. Instead of using the standard adapter and ports design, the Hexabus focuses on using events and plays to the natural strengths of Go's asynchronous go routines and channels. Generic events are sent to the bus and the callbacks occur via a channel included in the event itself. This yields automatic concurrency with ease of getting data back. 

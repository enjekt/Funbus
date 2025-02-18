# Funbus

The funbus is a bus that has a subsribe and unsubcribe mechanism for adding functions to specific topics that match the parameter of the function. The event name is a key to a map of map of fnName:channel. The channels are created automatically and the function name creates a unique value based on the what the reflection mechanism returns for the name (apparently the  address of the function.)

When an event is sent, reflection determines the name of the event and looks that up in the topic map. If a topic is found, the elements of that topic are iterated over and the event is passed onto the channel (stored as the value associated with a function "name").

The funbus is easy to use and easy to understand. No channels are required and there are only a few limitations (for example, the function must have a single event parameter).


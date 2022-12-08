/*
Datastore package define several interfaces for gchain to interact with data structure.

 1. Datastore
    It define a general interface to any database system available and use it to store and search for data.
 2. Vectorstore
    Fully compatible with datastore interface with addition of SearchVector functionality. Vector search database will prefer to use this interface.
 3. Retriever
    Provide simplified interface to get data from outside world, intended to interact in read only manner.

Datastore and V
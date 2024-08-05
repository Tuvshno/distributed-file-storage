# Distributed File Storage System
## Overview
This project is a decentralized, fully distributed content-addressable file storage system, developed entirely from scratch using Go. It leverages peer-to-peer (P2P) TCP communication to enable efficient and secure storage and retrieval of large files across a network of distributed nodes. The system is designed to handle large-scale data with an emphasis on reliability, scalability, and speed.

## Key Features
- Decentralized Architecture: Completely removes reliance on a central server, ensuring that the system is resilient, scalable, and free from single points of failure.
- Fully Distributed Content Addressing: Files are stored and retrieved based on their content, allowing for efficient deduplication and ensuring data integrity across the network.
- P2P TCP Communication: Implements a custom TCP library that facilitates robust and efficient peer-to-peer communication, enabling seamless data transmission between nodes.
- Streaming Large Files: Supports the streaming of large files across the network, optimizing bandwidth usage and reducing latency.
- Go Implementation: Entirely written in Go, taking advantage of its concurrency model and performance capabilities to handle the demands of a distributed system.
## What I Learned
Through the development of this project, I gained deep insights into the following areas:

- Distributed Systems: I explored the challenges and solutions associated with building distributed systems, such as network partitioning, consistency, and fault tolerance. This project solidified my understanding of decentralized architectures and the importance of designing systems that are both robust and scalable.

- Networking: Implementing a custom TCP library required a thorough understanding of low-level networking concepts, including socket programming, connection handling, and data serialization. I honed my skills in managing concurrent network connections, ensuring reliable and efficient data transfer.

- Content Addressable Storage: Building a content-addressable storage system involved learning about hashing algorithms and the design of systems that can efficiently manage and retrieve data based on content rather than location. This knowledge is critical for developing systems that require high data integrity and deduplication.

- Go: Writing the entire project in Go allowed me to deepen my expertise in this language, particularly in areas such as concurrency, memory management, and performance optimization. Go's strong typing and simplicity made it an ideal choice for implementing a complex system like this.

- File Streaming and Management: Developing the capability to stream large files across a distributed network taught me how to manage resources effectively, handle large datasets, and optimize for network bandwidth, ensuring smooth and efficient data transfers.

## Conclusion
This project really allowed me to design and implement complex distributed systems from the ground up, with a strong focus on network communication, data integrity, and performance. It sparked my interest into continuing to learn how to build more distributed systems from scratch and made me want to build my next project - a [distributed cache](https://github.com/Tuvshno/distributed-cache)
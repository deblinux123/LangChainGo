# 🦜 LangChainGo — 14-Day Learning Roadmap

A practical, 14-day roadmap for learning **LangChainGo**, the Go implementation of LangChain, and building AI-powered applications using Large Language Models (LLMs).

## 🎯 Goals

- Understand LangChainGo architecture and core components.
- Integrate LLMs with Go applications.
- Work with prompts, chains, and conversation memory.
- Implement embeddings and semantic search.
- Build Retrieval-Augmented Generation (RAG) applications.
- Explore agents, tools, and tool calling.
- Integrate LangChainGo with a REST API.
- Apply testing, error handling, and production best practices.

## 🛠️ Tech Stack

- **Language:** Go
- **LLM Framework:** LangChainGo
- **Local LLM Runtime:** Ollama
- **API Framework:** Fiber
- **Vector Database:** Qdrant
- **Cache and Session Storage:** Redis
- **Containerization:** Docker
- **Testing:** Go testing package and curl

## 📚 Learning Resources

- [Official Documentation](https://tmc.github.io/langchaingo/docs/)
- [Tutorials](https://tmc.github.io/langchaingo/docs/tutorials/)
- [How-to Guides](https://tmc.github.io/langchaingo/docs/how-to/)
- [Concepts and Architecture](https://tmc.github.io/langchaingo/docs/concepts/)
- [GitHub Repository](https://github.com/tmc/langchaingo)

---

## 📅 14-Day Plan

### Day 01 — Introduction to LangChainGo

**Topics**
- What is LangChainGo?
- Why use an LLM framework?
- LangChainGo architecture and core components.
- Understanding interfaces and modular design.
- Context propagation and error handling in Go.

**Practical Tasks**
- Read the introduction and architecture documentation.
- Create a Go project.
- Install LangChainGo.
- Explore the repository and its package structure.

**Deliverable:** A working Go project with LangChainGo installed.

### Day 02 — LLM Integration with Ollama

**Topics**
- LLM abstractions and model providers.
- Chat models versus text completion.
- Model configuration and generation parameters.
- Context, timeouts, and error handling.

**Practical Tasks**
- Install and run Ollama.
- Pull a compatible local model.
- Connect LangChainGo to Ollama.
- Send prompts and display model responses.
- Handle errors and request cancellation.

**Deliverable:** A command-line application that communicates with a local LLM.

### Day 03 — Prompts and Templates

**Topics**
- Prompt engineering fundamentals.
- Static and dynamic prompts.
- Prompt templates and variables.
- System, human, and AI messages.
- Few-shot prompting.

**Practical Tasks**
- Create reusable prompt templates.
- Build a text summarizer.
- Build a question-answering prompt.
- Experiment with different system instructions.
- Separate prompt construction from business logic.

**Deliverable:** A reusable prompt-based AI application.

### Day 04 — Chains and Composable Workflows

**Topics**
- What are chains?
- LLM chains and sequential workflows.
- Input and output keys.
- Composing multiple AI operations.
- Error propagation between workflow steps.

**Practical Tasks**
- Build a text summarization chain.
- Create a workflow that summarizes and categorizes text.
- Combine multiple processing steps.
- Test each component independently.

**Deliverable:** A multi-step LLM workflow implemented in Go.

### Day 05 — Conversation Memory

**Topics**
- Stateless versus stateful LLM requests.
- Conversation history.
- Buffer and window memory concepts.
- Context management.
- Persisting conversation state.

**Practical Tasks**
- Build a conversational chatbot.
- Maintain history across multiple messages.
- Implement session-based conversation handling.
- Explore persistent storage for conversation history.

**Deliverable:** A chatbot that maintains conversation context.

### Day 06 — Embeddings and Semantic Search

**Topics**
- What are embeddings?
- Semantic similarity and vector representations.
- Embedding model integration.
- Similarity search.
- Distance metrics and vector dimensions.

**Practical Tasks**
- Generate embeddings for sample documents.
- Compare semantically similar sentences.
- Implement a basic similarity search.
- Investigate embedding model requirements.

**Deliverable:** A semantic search application.

### Day 07 — Document Loading and Text Splitting

**Topics**
- Document loaders.
- Document metadata and schemas.
- Text splitting strategies.
- Chunk size and overlap.
- Preparing documents for retrieval.

**Practical Tasks**
- Load text files and Markdown documents.
- Split documents into chunks.
- Preserve source metadata.
- Experiment with chunk size and overlap.
- Prepare a small document collection for indexing.

**Deliverable:** A document ingestion pipeline.

### Day 08 — Vector Stores with Qdrant

**Topics**
- Vector store abstractions.
- Qdrant collections and points.
- Storing embeddings and metadata.
- Similarity search and filtering.
- Retriever interfaces.

**Practical Tasks**
- Run Qdrant with Docker.
- Configure a compatible embedding model.
- Index document chunks.
- Retrieve the most relevant documents.
- Test metadata filtering and retrieval quality.

**Deliverable:** A semantic document search system powered by Qdrant.

### Day 09 — Building a RAG Application

**Topics**
- Retrieval-Augmented Generation architecture.
- Retriever and LLM integration.
- Retrieval-based question answering.
- Context construction.
- Grounded responses and source attribution.

**Practical Tasks**
- Retrieve relevant document chunks.
- Construct prompts using retrieved context.
- Generate answers using a local LLM.
- Return source metadata with answers.
- Test questions that are and are not covered by the documents.

**Deliverable:** A working RAG application using Go, Ollama, and Qdrant.

### Day 10 — Agents and Tools

**Topics**
- Agent architecture and execution loops.
- Tools and tool interfaces.
- Tool selection and execution.
- Function calling.
- Multi-step workflows and error handling.

**Practical Tasks**
- Explore the available agent implementations.
- Implement or adapt a simple calculator tool.
- Build a tool-enabled assistant.
- Investigate the selected model's tool-calling compatibility.
- Handle tool failures and invalid inputs.

**Deliverable:** An AI assistant capable of invoking tools.

### Day 11 — Memory, Sessions, and Multi-Turn RAG

**Topics**
- Combining retrieval with conversation history.
- Session isolation.
- Short-term versus persistent memory.
- Managing conversation context.
- Avoiding irrelevant or stale context.

**Practical Tasks**
- Combine conversation memory with RAG.
- Use session IDs to separate conversations.
- Persist session history in Redis.
- Retrieve relevant documents for each user question.
- Test multiple independent sessions.

**Deliverable:** A session-aware RAG chatbot.

### Day 12 — Building a REST API with Fiber

**Topics**
- Integrating LangChainGo into a backend.
- Handler, service, and repository layers.
- Request validation and JSON responses.
- Streaming AI responses.
- Authentication, rate limiting, and timeouts.

**Practical Tasks**
- Build a Fiber REST API.
- Implement chat and health-check endpoints.
- Integrate the LLM and retrieval services.
- Add Redis-backed session handling.
- Implement consistent error responses.
- Test endpoints using curl.

**Deliverable:** A LangChainGo-powered AI REST API.

### Day 13 — Testing, Reliability, and Docker

**Topics**
- Unit and integration testing.
- Mocking model interactions.
- Logging and observability.
- Retry strategies and timeouts.
- Configuration and secret management.
- Docker-based deployment.

**Practical Tasks**
- Test prompts, services, and API handlers.
- Mock LLM responses where appropriate.
- Validate malformed requests and upstream failures.
- Add structured logging.
- Containerize the application.
- Verify service connectivity and health checks.

**Deliverable:** A tested, containerized AI backend.

### Day 14 — Final Project and Review

**Project: AI Knowledge Assistant**

Build a complete AI assistant using the concepts covered throughout the roadmap.

**Features**
- Local LLM integration with Ollama.
- Prompt templates and composable workflows.
- Document ingestion and chunking.
- Embeddings and semantic retrieval.
- Qdrant-backed RAG.
- Session-based conversation memory in Redis.
- Optional tool integration.
- Fiber REST API.
- Streaming responses where supported.
- Error handling, logging, and Docker deployment.

**Final Tasks**
- Review the official documentation.
- Refactor the project into modular packages.
- Write a README with setup and usage instructions.
- Add example API requests.
- Document architecture and configuration.
- Identify limitations and future improvements.

**Deliverable:** A documented, modular AI knowledge assistant built with Go.

---

## 📊 Progress Tracker

- [ ] Day 01 — Introduction to LangChainGo
- [ ] Day 02 — LLM Integration with Ollama
- [ ] Day 03 — Prompts and Templates
- [ ] Day 04 — Chains and Composable Workflows
- [ ] Day 05 — Conversation Memory
- [ ] Day 06 — Embeddings and Semantic Search
- [ ] Day 07 — Document Loading and Text Splitting
- [ ] Day 08 — Vector Stores with Qdrant
- [ ] Day 09 — Building a RAG Application
- [ ] Day 10 — Agents and Tools
- [ ] Day 11 — Memory, Sessions, and Multi-Turn RAG
- [ ] Day 12 — Building a REST API with Fiber
- [ ] Day 13 — Testing, Reliability, and Docker
- [ ] Day 14 — Final Project and Review

## 🏗️ Proposed Project Structure

```text
langchaingo-learning/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── services/
│   │   ├── llm/
│   │   ├── rag/
│   │   ├── memory/
│   │   └── tools/
│   └── repositories/
├── docs/
├── examples/
├── tests/
├── .env.example
├── .gitignore
├── go.mod
└── README.md
```

This is a suggested structure for the final project, not a requirement for every daily exercise.

## ⏱️ Daily Study Routine

Suggested daily schedule:

- **30 minutes:** Read the relevant official documentation.
- **60 minutes:** Implement the day's practical tasks.
- **20 minutes:** Test, debug, and review the code.
- **10 minutes:** Write notes and update the progress tracker.

## ✅ Learning Principles

1. Write code every day instead of only reading documentation.
2. Prefer small, independently testable components.
3. Understand the underlying Go interfaces rather than treating the framework as a black box.
4. Verify provider compatibility before selecting model features.
5. Keep secrets in environment variables and never commit credentials.
6. Check the current documentation and examples when APIs change.
7. Finish with a working project, not just a collection of tutorials.

## 🎓 Expected Outcome

By the end of these 14 days, the goal is to understand the core LangChainGo abstractions and apply them to a practical AI backend with local inference, retrieval, conversation memory, and REST API integration.

**Important:** Fourteen days is an introduction and implementation sprint, not mastery of every feature. Continue exploring the official API reference, examples, and production patterns after completing the roadmap.

---

**Learning resource:** [LangChainGo Official Documentation](https://tmc.github.io/langchaingo/docs/)

**Repository:** [tmc/langchaingo](https://github.com/tmc/langchaingo)
> [!IMPORTANT]
> **Chibi — Final Release Documentation (2026)**
> This documentation represents the final public release for GitHub. Stable / Open Source / 2026 Edition.

## Core Components

This section provides a detailed description of Chibi's core components, outlining their functionalities, interactions, and technical considerations. These components are the building blocks that enable Chibi's intelligent and efficient Kubernetes SRE capabilities.

### 1. AI System

The AI System is the brain of Chibi, responsible for understanding user intent, diagnosing issues, generating solutions, and orchestrating interactions with various AI models. It is designed to be modular, extensible, and context-aware.

#### 1.1. AI Abstraction Layer

*   **Functionality:** Provides a unified interface for Chibi to interact with diverse AI providers (e.g., OpenAI, Claude, Gemini, local models). It decouples the core logic from specific AI model implementations, allowing for seamless integration of new models and dynamic selection.
*   **Key Features:**
    *   **Standardized API:** Presents a consistent API to the rest of Chibi, regardless of the underlying AI provider.
    *   **Request/Response Normalization:** Handles differences in input/output formats across various AI models.
    *   **Error Handling & Fallback:** Manages errors from AI providers and implements fallback strategies to ensure continuous operation.

#### 1.2. Provider Router

*   **Functionality:** Intelligently directs AI requests to the most suitable AI provider based on a set of dynamic criteria. This ensures optimal performance, cost-efficiency, and accuracy for each specific SRE task.
*   **Key Features:**
    *   **Contextual Routing:** Analyzes the nature of the request (e.g., complexity, sensitivity, urgency) to determine the best provider.
    *   **Cost Optimization:** Prioritizes cost-effective models for simpler tasks, reserving more expensive models for complex reasoning.
    *   **Latency Awareness:** Routes requests to providers with lower latency for time-sensitive operations.
    *   **Capability Matching:** Selects providers based on their specific strengths (e.g., code generation, natural language understanding, summarization).
    *   **Load Balancing:** Distributes requests across multiple providers to prevent rate limiting and ensure high availability.

#### 1.3. Context Builder

*   **Functionality:** Gathers and synthesizes relevant information to enrich AI prompts, ensuring that AI responses are highly contextual and accurate. This includes real-time cluster state, historical interactions, user preferences, and relevant documentation.
*   **Key Features:**
    *   **Kubernetes Context:** Injects information about the current cluster, namespace, selected resources, and their states.
    *   **Session History:** Incorporates previous commands and AI interactions to maintain conversational coherence.
    *   **User Preferences:** Integrates user-defined settings and preferences into prompts.
    *   **Documentation Integration:** Can pull relevant documentation snippets (e.g., Kubernetes API docs, best practices) to provide to the AI.
    *   **Token Optimization:** Summarizes and condenses context to fit within AI model token limits without losing critical information.

#### 1.4. Reasoning Engine

*   **Functionality:** The core intelligence component that processes contextualized prompts, performs complex analysis, and formulates intelligent responses or recommendations. It leverages advanced AI capabilities for diagnostics, predictive analysis, and solution generation.
*   **Key Features:**
    *   **Natural Language Understanding (NLU):** Interprets user commands and queries in natural language.
    *   **Problem Diagnosis:** Analyzes logs, metrics, and events to identify root causes of Kubernetes issues.
    *   **Predictive Analysis:** Identifies patterns and trends to predict potential incidents before they occur.
    *   **Solution Generation:** Proposes actionable remediation steps, configuration changes, or new resource manifests.
    *   **Explanation Generation:** Provides clear, human-readable explanations for diagnoses and proposed solutions.
    *   **Safety & Guardrails:** Implements mechanisms to prevent AI hallucinations and ensure that generated actions are safe and valid.

#### 1.5. Prompt Pipeline

*   **Functionality:** Responsible for constructing, optimizing, and managing prompts sent to AI providers. It ensures that prompts are well-formed, context-rich, token-efficient, and tailored to elicit the most accurate and relevant responses.
*   **Key Features:**
    *   **Prompt Templating:** Uses predefined templates to structure prompts for different SRE tasks.
    *   **Context Injection:** Dynamically inserts context from the Context Builder into the prompt.
    *   **Token Management:** Monitors and optimizes prompt length to stay within AI model token limits.
    *   **Instruction Tuning:** Adjusts prompt instructions based on the selected AI provider and task type.
    *   **Response Parsing Instructions:** Includes instructions for the AI on how to format its response for easier parsing by Chibi.

### 2. Kubernetes Engine

The Kubernetes Engine is Chibi's dedicated interface for interacting with Kubernetes clusters. It abstracts the complexities of the Kubernetes API, providing a reliable and efficient way to manage and observe cluster resources.

#### 2.1. `client-go` Integration

*   **Functionality:** Utilizes the official `client-go` library for Go applications to communicate with the Kubernetes API server. This ensures robust, idiomatic, and up-to-date interaction with Kubernetes.
*   **Key Features:**
    *   **API Client:** Provides programmatic access to all Kubernetes API resources.
    *   **Informers & Listers:** Efficiently watches for changes in Kubernetes resources and maintains local caches for fast read access.
    *   **Authentication & Authorization:** Inherits `kubeconfig` settings for secure authentication and respects Kubernetes RBAC policies.

#### 2.2. Resource Management

*   **Functionality:** Enables Chibi to perform CRUD (Create, Read, Update, Delete) operations on Kubernetes resources, driven by user commands or AI-generated actions.
*   **Key Features:**
    *   **Declarative Management:** Supports applying and managing resources using YAML/JSON manifests.
    *   **Imperative Commands:** Executes direct commands for common operations (e.g., scaling deployments, deleting pods).
    *   **Validation:** Validates resource definitions against Kubernetes API schemas before application.

#### 2.3. Data Retrieval & Streaming

*   **Functionality:** Collects real-time and historical data from Kubernetes clusters, including logs, events, metrics, and resource states, essential for diagnostics and visualizations.
*   **Key Features:**
    *   **Log Streaming:** Provides real-time streaming of container logs.
    *   **Event Monitoring:** Monitors and filters Kubernetes events for anomalies and state changes.
    *   **Metrics Integration:** Integrates with Prometheus or other metrics systems to retrieve resource utilization data.
    *   **Resource Watchers:** Continuously monitors specific resources for changes, enabling real-time updates in the UI.

#### 2.4. Multi-Cluster Context Management

*   **Functionality:** Allows users to seamlessly switch between and manage multiple Kubernetes clusters from a single Chibi instance. It maintains separate contexts and configurations for each cluster.
*   **Key Features:**
    *   **Context Switching:** Provides commands to easily switch the active Kubernetes context.
    *   **Unified View:** Offers aggregated views of resources and events across multiple configured clusters.
    *   **Configuration Storage:** Securely stores and manages `kubeconfig` files and cluster-specific settings.

### 3. CLI User Experience (CLI UX)

The CLI UX component focuses on creating an intuitive, efficient, and visually engaging terminal experience. It combines robust command-line parsing with rich, interactive UI elements and intelligent feedback.

#### 3.1. Command Parser & Autocompletion

*   **Functionality:** Interprets user input, parses commands and arguments, and provides intelligent autocompletion suggestions to enhance productivity and reduce errors.
*   **Key Features:**
    *   **Natural Language Processing (NLP):** Integrates with the AI System to understand natural language commands.
    *   **Contextual Autocompletion:** Suggests Kubernetes resources, flags, and values based on the current command and cluster state.
    *   **Error Correction:** Offers suggestions for misspelled commands or arguments.

#### 3.2. Interactive Terminal UI (TUI)

*   **Functionality:** Leverages modern TUI frameworks (e.g., Bubble Tea, Lip Gloss, Bubbles) to create a dynamic and responsive user interface directly within the terminal, going beyond traditional static text output.
*   **Key Features:**
    *   **Dynamic Layouts:** Adapts the display based on content and terminal size.
    *   **Interactive Widgets:** Provides interactive elements like lists, tables, forms, and progress bars.
    *   **Keyboard Navigation:** Optimized for keyboard-driven interaction for speed and efficiency.
    *   **Theming & Styling:** Supports customizable themes and rich text styling (colors, bold, italics).

#### 3.3. Terminal Rendering Engine

*   **Functionality:** Responsible for generating rich, interactive visualizations (e.g., Mermaid diagrams, live resource graphs, dependency trees) directly within the terminal, updating in real-time based on AI analysis and user interaction.
*   **Key Features:**
    *   **Mermaid Integration:** Renders Mermaid diagrams directly in the terminal for architectural views and workflows.
    *   **ASCII/Unicode Graphics:** Generates text-based graphs, charts, and visual representations of cluster state.
    *   **Real-time Updates:** Dynamically refreshes visualizations as cluster state or data changes.
    *   **Interactive Elements:** Allows users to interact with visualizations (e.g., hover for details, click to drill down) using keyboard shortcuts.

#### 3.4. Feedback & Explanations

*   **Functionality:** Provides clear, concise, and contextual feedback to the user, including explanations of AI decisions, diagnostic results, and command outcomes.
*   **Key Features:**
    *   **AI-Generated Explanations:** Presents AI's reasoning for diagnoses, predictions, and proposed remediations.
    *   **Actionable Error Messages:** Provides user-friendly error messages with clear steps for resolution.
    *   **Progress Indicators:** Shows the status of long-running operations.
    *   **Contextual Help:** Offers relevant help and documentation snippets based on the current command or context.

### 4. Plugins

The Plugin System is designed to make Chibi highly extensible, allowing the community and enterprises to add new functionalities, integrate with external tools, and customize behavior without modifying the core codebase.

#### 4.1. Plugin SDK (Software Development Kit)

*   **Functionality:** Provides a comprehensive set of tools, libraries, and documentation for developers to create new plugins for Chibi. It defines the interfaces and conventions for plugin development.
*   **Key Features:**
    *   **Go-based API:** A well-defined Go API for interacting with Chibi's core components (AI, Kubernetes, CLI UX).
    *   **Extension Points:** Clearly defined interfaces for adding new commands, subcommands, AI providers, data sources, and visualization types.
    *   **Helper Utilities:** Libraries for common tasks like configuration management, logging, and error handling within plugins.
    *   **Documentation & Examples:** Comprehensive guides and sample plugins to facilitate development.

#### 4.2. Plugin Runtime & Lifecycle

*   **Functionality:** Manages the loading, execution, and unloading of plugins, ensuring stability, security, and compatibility with the core Chibi application.
*   **Key Features:**
    *   **Dynamic Loading:** Loads plugins at runtime from specified directories or remote sources.
    *   **Version Management:** Ensures compatibility between plugins and the Chibi core, handling version conflicts.
    *   **Hot Reloading:** Allows for updating plugins without restarting Chibi (for development/testing).
    *   **Resource Management:** Monitors and manages resources consumed by plugins.

#### 4.3. Plugin Sandboxing & Security

*   **Functionality:** Provides a secure execution environment for plugins, isolating them from the core Chibi process and preventing unauthorized access to system resources or sensitive data.
*   **Key Features:**
    *   **Process Isolation:** Runs plugins in separate processes or sandboxed environments (e.g., WebAssembly) to limit their impact on the core.
    *   **Permission Model:** Defines a granular permission model for plugins, controlling their access to Kubernetes API, filesystem, network, and AI providers.
    *   **Code Signing & Verification:** Supports verifying the authenticity and integrity of plugins before loading them.
    *   **Auditing:** Logs plugin actions and resource usage for security auditing.

#### 4.4. Plugin Discovery & Installation

*   **Functionality:** Facilitates the discovery and installation of plugins by users, making it easy to extend Chibi's capabilities.
*   **Key Features:**
    *   **Local Plugin Management:** Commands for listing, installing, updating, and uninstalling plugins from local paths.
    *   **Remote Registry (Future):** A potential future feature for a centralized, curated registry of trusted Chibi plugins.
    *   **Dependency Management:** Handles plugin dependencies to ensure all required components are present.

These core components, working in concert, form the foundation of Chibi, enabling its unique blend of AI intelligence, Kubernetes expertise, and terminal-native user experience.


## References

[1] Kubernetes. (n.d.). *kubectl Overview*. Retrieved from [https://kubernetes.io/docs/reference/kubectl/overview/](https://kubernetes.io/docs/reference/kubectl/overview/)
[2] k9s. (n.d.). *k9s - Kubernetes CLI To Manage Your Clusters In Style!*. Retrieved from [https://k9scli.io/](https://k9scli.io/)
[3] Lens. (n.d.). *The Kubernetes IDE*. Retrieved from [https://k8slens.dev/](https://k8slens.dev/)
[4] Headlamp. (n.d.). *An extensible Kubernetes UI*. Retrieved from [https://headlamp.dev/](https://headlamp.dev/)
[5] Radar. (n.d.). *Radar - The Kubernetes UI*. Retrieved from [https://radarhq.io/](https://radarhq.io/)
[6] SRExpert. (n.d.). *Unified Kubernetes Management Platform*. Retrieved from [https://srexpert.io/](https://srexpert.io/)
[7] Kubeshark. (n.d.). *API Traffic Viewer for Kubernetes*. Retrieved from [https://kubeshark.co/](https://kubeshark.co/)
[8] Komodor. (n.d.). *Autonomous AI SRE Platform for Kubernetes*. Retrieved from [https://komodor.com/](https://komodor.com/)
[9] Kubecost. (n.d.). *Kubernetes Cost Monitoring & Optimization*. Retrieved from [https://www.kubecost.com/](https://www.kubecost.com/)
[10] Devtron. (n.d.). *AI-Native Kubernetes Management Platform*. Retrieved from [https://devtron.ai/](https://devtron.ai/)
[11] Portainer. (n.d.). *Universal Container Management Platform*. Retrieved from [https://www.portainer.io/](https://www.portainer.io/)
[12] Aider. (n.d.). *AI pair programming in your terminal*. Retrieved from [https://aider.chat/](https://aider.chat/)
[13] OpenCode. (n.d.). *Open-source AI coding agent*. Retrieved from [https://github.com/opencode/opencode](https://github.com/opencode/opencode)


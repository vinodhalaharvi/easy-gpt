# Easy-GPT Project

The Easy-GPT project provides two separate functionalities: a chat interface and an assistant interface, both leveraging
the OpenAI API. This README outlines how to build these interfaces, configure your environment, and optionally set up
Vim shortcuts for convenience.
Prerequisites

    Go (1.21 tested, other versions may work)
    An OpenAI API key

# Installation

This project serves as an AI wrapper over Linux commands, Kubernetes (kubectl), and Google Cloud (gcloud) operations,
enhancing the command-line experience with the power of AI.
It provides a conversational interface to interact with these systems more intuitively. However, it's important to note
that this project does not replace the underlying tools.
Users must have the necessary Linux, Kubernetes, and Google Cloud command-line tools installed and properly configured
on their system to use this project effectively.

# Configuration

Before running the chat and assistant, you need to configure your environment:

Export the OpenAI API Key: Set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY='your_openai_api_key_here'
```

Set the Configuration File Path: Point to your config.yaml file using an environment variable:

```bash
export CONFIG_FILE_PATH='./config.yaml'
```

Sample config looks like this. Edit if you want to change model. But everything else should remain same.

```text
chat:
  model: "gpt-4-turbo-preview"
  requestMethod: "POST"
  requestURL: "https://api.openai.com/v1/chat/completions"
  messages:
    - role: "system"
      content: |
        You are a versatile Linux assistant familiar with a wide range of commands
        including but not limited to general Linux system commands, Google Cloud (gcloud), and Kubernetes (kubectl) operations.
        Focus on providing the most direct command or commands to achieve the user's query.

assistant:
  model: "gpt-4-turbo-preview"
  systemMessage: |
    You are a versatile Linux assistant familiar with a wide range of commands
    including but not limited to general Linux system commands, Google Cloud (gcloud), and Kubernetes (kubectl) operations.
    Respond with a JSON array of commands without explanations, Markdown, or additional formatting.
    Provide the response in plain JSON format. Focus on providing the most direct command or commands to achieve the user's query.
  requestMethod: "POST"
  requestURL: "https://api.openai.com/v1/chat/completions"
```

Ensure your config.yaml is correctly formatted and located at the specified path.

# Building the Binaries

Navigate to the root directory of the Easy-GPT project to build both the chat and assistant binaries:

# Build the Assistant:

```bash
cd cmd/assistant
go build -o gpt-assistant
cd ../..
```

Build the Chat:

```bash
cd cmd/chat
go build -o gpt-chat
cd ../..
```

This process will create two executables: gpt-assistant and gpt-chat, in their respective directories.
Running the Binaries

After building, you can run each binary directly from the command line:

For the Assistant:

```bash
echo list all namespaces | ./cmd/assistant/gpt-assistant
```

For the Chat:

```bash
complete this ip range 127.0.0.1-15 | ./cmd/chat/gpt-chat 
```

# Optional: Creating Vim Shortcuts

If you frequently use Vim and wish to integrate shortcuts for running the chat and assistant, you can add the following
to your .vimrc file:

```vim
" Assistant Shortcut
nnoremap <silent> K :w !./cmd/assistant/gpt-assistant<CR>

" Chat Shortcut
nnoremap <silent> M :w !./cmd/chat/gpt-chat<CR>

" Bash Shortcut
nnoremap <silent> B :w !bash<CR>

```

These mappings will allow you to press Shift+B, M, K to send the current buffer's content to the bash, chat or assistant
respectively.

## Voilà! Profit

You're all set! You can now interact with the Easy-GPT project's chat and assistant functionalities, leveraging the
power of OpenAI's GPT models right from your terminal or Vim editor.

# MIT License

# Like my work?

If you like my work and want to hire me for your Go project needs, you can reach me
at [vinod@smartify.software](mailto:vinod@smartify.software)
or [LinkedIn](https://www.linkedin.com/in/vinod-halaharvi-289a1a13/).

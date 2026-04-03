# 🧠 AI Agent Project - Boot.dev

Welcome to the AI Agent Project built as part of the [Boot.dev](https://boot.dev) course on building autonomous agents!

This project demonstrates a basic AI agent capable of perceiving, reasoning, and acting autonomously in a defined environment.

## 🚀 Project Features

* Autonomous task execution loop
* Memory management (short- and long-term)
* Dynamic prompt creation
* Integration with LLM APIs (Google Gemini)
* Command parsing and execution


## 📁 Project Structure

```
.
├── poetry.lock                        # Defines the project's locked dependencies with Poetry
├── pyproject.toml                     # The project configuration file
└── src                                 
    ├── calculator                     # The AI Agent wokrking directory (contains an example code repo) 
    ├── config.py                      # Contains project configuration variable defraults
    ├── function_schema.py             # The function schema definitions for the agent
    ├── functions                      # The functions available to the agent
    │   ├── call_function.py
    │   ├── get_file_content.py
    │   ├── get_files_info.py
    │   ├── run_python_file.py
    │   └── write_file.py
    ├── main.py                        # The main process for the project
    └── utils.py                       # contains util functions
```

## ⚙️ Requirements

* Python 3.8+
* Google Gemini API key
* Dependencies (install with pipx):

### This Project uses [Poetry](https://python-poetry.org/)
Poetry is a tool for dependency management and packaging in Python. It allows you to declare the libraries your project depends on and it will manage (install/update) them for you. Poetry offers a lockfile to ensure repeatable installs, and can build your project for distribution.

#### To get started
```bash
pipx install poetry
poetry install           # Installs project dependencies
poetry shell             # Creates a virtual envrionment to execute the project code

```

## 🔑 Setup

1. Go to Google AI Studio, create an account, and [get your API key](https://ai.google.dev/gemini-api/docs/api-key)


2. Create a .env in the project root and set your API key as an environment variable:

   ```bash
   export GEMINI_API_KEY=your-api-key-here
   ```

3. Run the agent:
    You can optionally run your requests with a `--verbose` flag for more details outputs

   ```bash
   python main.py "Your AI agent request" --verbose
   ```

## 📚 Learn More

This is a guided project from Boot.dev's AI Agent Python course. You can learn more and follow along at [boot.dev/courses](https://boot.dev/courses).

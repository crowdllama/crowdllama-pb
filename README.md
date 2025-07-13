# CrowdLlama Protobuf Definitions

This repository contains Protocol Buffer definitions for the CrowdLlama API, specifically for interacting with Llama models.

## Structure

The Protobuf definitions include:

### LlamaRequest
Represents a request to the Llama API:
- `model`: The model to use (e.g., "llama3.2")
- `prompt`: The input prompt
- `stream`: Whether to stream the response

### LlamaResponse
Represents a response from the Llama API:
- `model`: The model used
- `created_at`: Timestamp when response was created
- `response`: The generated response text
- `done`: Whether the response is complete
- `done_reason`: Reason for completion (e.g., "stop")
- `context`: Token context array
- `total_duration`: Total duration in nanoseconds
- `load_duration`: Model loading duration in nanoseconds
- `prompt_eval_count`: Number of prompt tokens evaluated
- `prompt_eval_duration`: Prompt evaluation duration in nanoseconds
- `eval_count`: Number of tokens evaluated
- `eval_duration`: Evaluation duration in nanoseconds

## Service Definition

The `LlamaService` provides two RPC methods:
- `Generate`: Generates a single response
- `GenerateStream`: Generates a streaming response

## Usage

### Generate Go code

To generate Go code from the Protobuf definitions:

```bash
# Install buf if you haven't already
go install github.com/bufbuild/buf/cmd/buf@latest

# Generate Go code
buf generate
```

### Example JSON to Protobuf

**Request:**
```json
{
  "model": "llama3.2",
  "prompt": "1+1",
  "stream": false
}
```

**Response:**
```json
{
  "model": "llama3.2",
  "created_at": "2025-07-13T05:29:25.342493Z",
  "response": "1 + 1 = 2",
  "done": true,
  "done_reason": "stop",
  "context": [128006, 9125, 128007, 271, 38766, 1303, 33025, 2696, 25, 6790, 220, 2366, 18, 271, 128009, 128006, 882, 128007, 271, 16, 10, 16, 128009, 128006, 78191, 128007, 271, 16, 489, 220, 16, 284, 220, 17],
  "total_duration": 299529791,
  "load_duration": 54093833,
  "prompt_eval_count": 28,
  "prompt_eval_duration": 156589417,
  "eval_count": 8,
  "eval_duration": 88340208
}
```

## Development

This project uses [Buf](https://buf.build/) for Protobuf development. The configuration files include:

- `buf.yaml`: Main configuration
- `buf.gen.yaml`: Code generation configuration
- `go.mod`: Go module definition

### CI/CD

The project includes GitHub Actions workflows that automatically:

- **Format checking**: Ensures Protobuf files are properly formatted
- **Linting**: Validates Protobuf syntax and style
- **Breaking changes**: Detects API breaking changes
- **Code generation**: Ensures generated Go code is up to date

### Pre-commit Hooks

Install pre-commit hooks to catch issues before committing:

```bash
# Install pre-commit
pip install pre-commit

# Install the git hook scripts
pre-commit install
```

This will automatically run Protobuf linting and formatting on every commit.

## License

See [LICENSE](LICENSE) file for details. 
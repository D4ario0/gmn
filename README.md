# gmn
Dead Simple Terminal Gemini Client to stop opening the browser

# Features
- Gemini API.
- Syntax Highlithing.
- Web-like result output.
- No config required.
- Multiline input.

# Install
```bash
go install gihtub.com/D4ario0/gmn
```

Make sure to add `GOOGLE_API_KEY` in your `.bashrc` or `config.fish` a GEMINI API Key from Google AI Studio.
```bash
export GOOGLE_API_KEY="<your-key>"
```

# Usage
This works like an AI Shell like Warp but without the extras.

```bash
gmn
```

Type your prompt. Press `Ctrl + d` to send prompt. `Ctrl + c` to exit.

export const Prompt = (gitDiff: string) => `
You are an expert at writing Git commit titles.

Your ONLY task is to write **one short, single-line Git commit title**.

### RULES (follow these exactly)
1. Output ONLY the commit title — absolutely **no explanations, no formatting, no markdown, no punctuation around it**.
2. Do NOT use quotes (" "), backticks (\` \`), or any kind of brackets.
3. Use **Conventional Commit** format (e.g., feat:, fix:, refactor:, chore:, docs:, style:, etc.).
4. Write in a **formal** tone.
5. Keep it **short** (under ~10 words).
6. Do NOT include file names, code, variables, or underscores.
7. The result must look like this (example):  
   feat: improve user authentication flow  
   ✅ Correct  
   ❌ "feat: improve user authentication flow"  
   ❌ \`feat: improve user authentication flow\`  
   ❌ feat: improve user authentication flow - updated user.go

### Configuration
- Conventional commits: true
- Use emojis: false
- Tone: formal
- Length: short

### Input
git diff:
${gitDiff}

### Output
Now write ONLY the Git commit title — one single line, plain text, no quotes or formatting:`;

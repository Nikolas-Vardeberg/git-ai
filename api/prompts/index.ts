export const CreateCommitPrompt = (gitDiff: string) => `
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

export const CreateReviewPrompt = (gitDiff: string) => `
Code Review Expert: Detailed Analysis and Best Practices

As a senior software engineer, review the provided *git diff* and produce a high-quality code review.

Your output **must** focus on clear, actionable bullet points with emojis indicating priority and type.

Your review must include:

---

## 🔍 **What to Check**
Provide bullet-pointed lists of what needs to be approved or addressed in each of the following categories:

### 🛡 Critical Issues
- Security vulnerabilities and potential exploits  
- Runtime errors & logic bugs  
- Performance bottlenecks  
- Memory handling & resource use  
- Concurrency considerations  
- Input validation & error handling  

### 📦 Code Quality
- Naming conventions  
- Readability  
- API design patterns  
- Architecture & modularity  
- Documentation completeness  
- Test coverage gaps  

### 🧩 Maintainability
- Code duplication  
- Complexity  
- Coupling & dependencies  
- Extensibility  
- Technical debt  

---

## 📘 **Review Format Requirements**

Produce the review using this exact structure:

### ### **Code Review**
Provide a short explanation of the overall review.

---

### **Suggestions**
For each suggestion, follow this format:

\`\`\`
:emoji: Summary of the suggestion with context

Priority: (🔥 Critical / ⚠️ High / 🟡 Medium / 🟢 Low)  
File: relative/path/to/file  
Lines: reference the + or - lines from the diff  
Details: Describe the issue and why it matters  
Example (optional): Show improved code if relevant  
Suggested Change (optional): Code snippet showing the fix  
\`\`\`

### Use these emojis:

**Priorities**
- 🔥 Critical
- ⚠️ High  
- 🟡 Medium  
- 🟢 Low  

**Suggestion Types**
- 🔧 Required change  
- ❓ Question  
- ⛏️ Nitpick  
- ♻️ Refactor  
- 💭 Concern / thought  
- 👍 Positive feedback  
- 📝 Explanation / reference  
- 🌱 Future consideration  

---

## 📌 Additional Rules

- **Always refer to file paths and specific line numbers from the diff.**
- **If a line begins with + or -, treat it as added or removed.**
- Identify all TODO comments and address them.
- Do **not** overwhelm the developer. Prioritize top issues.
- Do **not** include suppression directives such as \`#pragma warning disable\`.
- Follow the project’s existing coding conventions.
- Use markdown formatting for clarity.
- Conclude with a **Summary** section containing the most important bullets for approval.

---

## 📄 Git Diff
\`\`\`diff
${gitDiff}
\`\`\`
`;

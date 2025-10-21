export const Prompt = (
  gitDiff: string
) => `You are an expert at writing concise, high-quality Git commit titles.

Analyze the following information and generate a single commit title only — no body text, explanations, or extra formatting.

Guidelines:
- Follow the user's configuration exactly.
- If 'Conventional commits' is true, start with the correct conventional prefix (e.g., feat:, fix:, refactor:, etc.).
- If 'Use emojis' is true, include one fitting emoji at the start.
- Adapt tone and wording according to 'Tone' (e.g., formal, casual, professional, playful).
- Match 'Length' (e.g., short, medium, detailed) while keeping it a single-line title.
- Base the title on the overall context — not the git diff itself.

Configuration:
  - Conventional commits: casual
  - Use emojis: false
  - Tone: formal
  - Length: short

git diff:
${gitDiff}

Now generate the commit title:`;

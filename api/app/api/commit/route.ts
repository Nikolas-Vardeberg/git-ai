/* eslint-disable @typescript-eslint/no-explicit-any */
import { NextRequest, NextResponse } from "next/server";
import { CreateCommitPrompt } from "@/prompts";
import { generateText } from "ai";
import { createGoogleGenerativeAI } from "@ai-sdk/google";

const google = createGoogleGenerativeAI();

export async function POST(req: NextRequest) {
  try {
    const { gitDiff } = await req.json();

    if (!gitDiff) {
      return NextResponse.json({ status: 400, error: "Missing gitDiff" });
    }

    const { text } = await generateText({
      model: google("gemini-2.5-flash"),
      prompt: CreateCommitPrompt(gitDiff),
    });

    const commitMessage = text || "No commit message generated";

    return NextResponse.json({
      data: { commitMessage },
    });
  } catch (error: any) {
    NextResponse.json({ error });
  }
}

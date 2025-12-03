import { createGoogleGenerativeAI } from "@ai-sdk/google";
import { NextRequest, NextResponse } from "next/server";
import { generateText } from "ai";
import { CreateReviewPrompt } from "@/prompts";

const google = createGoogleGenerativeAI();

export async function POST(req: NextRequest) {
  try {
    const { gitDiff } = await req.json();

    if (!gitDiff) {
      return NextResponse.json({ status: 400, error: "Missing gitDiff" });
    }

    const { text } = await generateText({
      model: google("gemini-2.5-flash"),
      prompt: CreateReviewPrompt(gitDiff),
    });

    if (!text) {
      return NextResponse.json({ status: 400, error: "Failed to generate" });
    }

    return NextResponse.json({ text });
  } catch (error) {
    console.log(error);
  }
}

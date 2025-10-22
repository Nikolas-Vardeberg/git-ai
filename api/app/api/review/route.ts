/* eslint-disable @typescript-eslint/no-explicit-any */
import { NextRequest, NextResponse } from "next/server";
import axios from "axios";
import { CreateReviewPrompt } from "@/prompts";

export async function POST(req: NextRequest) {
  try {
    const { gitDiff } = await req.json();

    if (!gitDiff) {
      return NextResponse.json({ status: 400, error: "Missing gitDiff" });
    }

    const response = await axios.post(process.env.OLLAMA_SERVER!, {
      model: "mistral:instruct",
      prompt: CreateReviewPrompt(gitDiff),
      stream: false,
    });

    const reviewMessage =
      response?.data?.response || "No commit message generated";

    return NextResponse.json({
      data: { reviewMessage },
    });
  } catch (error: any) {
    NextResponse.json({ error });
  }
}

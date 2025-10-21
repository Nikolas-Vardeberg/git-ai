import { NextRequest, NextResponse } from "next/server";
import axios from "axios";
import { Prompt } from "@/prompts";

export async function POST(req: NextRequest) {
  try {
    const { gitDiff } = await req.json();

    if (!gitDiff) {
      return NextResponse.json({ status: 400, error: "Missing gitDiff" });
    }

    const response = await axios.post("http://127.0.0.1:11434/api/generate", {
      model: "mistral:instruct",
      prompt: Prompt(gitDiff),
      stream: false,
    });

    const commitMessage =
      response?.data?.response || "No commit message generated";

    return NextResponse.json({
      data: { commitMessage },
    });
  } catch (error: any) {
    console.error("API /commit error:", error.message);
    return NextResponse.json({
      status: 500,
      error: error.message || "Internal Server Error",
    });
  }
}

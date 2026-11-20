// Written by apm install. Pi has no command hooks, so this extension hands every finished tool
// call to `apm hook pi` and appends what it prints to the result the model sees next.
import { execFileSync } from "node:child_process";
import type { ExtensionAPI } from "@earendil-works/pi-coding-agent";

export default function apm(pi: ExtensionAPI) {
	pi.on("tool_result", async (event, ctx) => {
		const payload = JSON.stringify({
			session_id: ctx.sessionManager.getSessionId(),
			cwd: ctx.cwd,
			hook_event_name: "tool_result",
			tool_name: event.toolName,
			tool_input: event.input,
		});
		let text = "";
		try {
			text = execFileSync("apm", ["hook", "pi"], { input: payload, encoding: "utf8", timeout: 5000 });
		} catch (error) {
			if ((error as NodeJS.ErrnoException).code === "ENOENT") return; // apm is not installed here
			ctx.ui.notify(`apm hook failed: ${error instanceof Error ? error.message : String(error)}`, "error");
			return;
		}
		if (text.trim() === "") return;
		return { content: [...event.content, { type: "text", text }] };
	});

	pi.on("session_compact", async (_event, ctx) => {
		try {
			execFileSync("apm", ["hook", "pi"], {
				input: JSON.stringify({ session_id: ctx.sessionManager.getSessionId(), cwd: ctx.cwd, hook_event_name: "compact" }),
				timeout: 5000,
			});
		} catch {}
	});
}

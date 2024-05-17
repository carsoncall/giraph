// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import { getWebviewOptions, GiraphPanel } from "./GiraphPanel.js";

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  context.subscriptions.push(
    vscode.commands.registerCommand("giraph.start", () => {
      GiraphPanel.createOrShow(context.extensionUri);
    })
  );

  context.subscriptions.push(
    vscode.commands.registerCommand("giraph.doRefactor", () => {
      if (GiraphPanel.currentPanel) {
        GiraphPanel.currentPanel.doRefactor();
      }
    })
  );

  if (vscode.window.registerWebviewPanelSerializer) {
    // make sure we register a serializer in activation event
    vscode.window.registerWebviewPanelSerializer(GiraphPanel.viewType, {
      async deserializeWebviewPanel(
        webviewPanel: vscode.WebviewPanel,
        state: any
      ) {
        console.log(`Got state: ${state}`);
        webviewPanel.webview.options = getWebviewOptions(context.extensionUri);
        GiraphPanel.revive(webviewPanel, context.extensionUri);
      },
    });
  }
}

// This method is called when your extension is deactivated
export function deactivate() {}

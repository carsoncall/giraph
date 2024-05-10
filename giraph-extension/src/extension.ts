// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  vscode.commands.registerCommand("giraph.start", () => {
    // Create and show a new webview
    const panel = vscode.window.createWebviewPanel(
      "giraph", // Identifies the type of the webview. Used internally
      "Giraph", // Title of the panel displayed to the user
      vscode.ViewColumn.One, // Editor column to show the new webview panel in.
      {} // Webview options. More on these later.
    );

    panel.webview.html = getWebviewContent();
  });

  let disposable = vscode.commands.registerCommand("giraph.helloWorld", () => {
    // The code you place here will be executed every time your command is executed
    // Display a message box to the user
    vscode.window.showInformationMessage("Hello World from Giraph!");
  });

  context.subscriptions.push(disposable);
}

// This method is called when your extension is deactivated
export function deactivate() {}

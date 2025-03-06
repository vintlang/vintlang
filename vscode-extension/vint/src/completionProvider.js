const vscode = require("vscode");

function activate(context) {
  let provider = vscode.languages.registerCompletionItemProvider("vint", {
    provideCompletionItems(document, position, token, context) {
      const keywords = [
        "func", "return", "if", "else", "while", "for", "switch", "case", "default", "break", "continue"
      ];
      return keywords.map(keyword => {
        let item = new vscode.CompletionItem(keyword, vscode.CompletionItemKind.Keyword);
        return item;
      });
    }
  });

  context.subscriptions.push(provider);
}

exports.activate = activate;

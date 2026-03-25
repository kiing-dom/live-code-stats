import * as vscode from 'vscode'

export function activate(context: vscode.ExtensionContext) {

    vscode.workspace.onDidChangeTextDocument(() => {
        sendUpdate({ keystrokes: 1});
    });

    vscode.workspace.onDidSaveTextDocument((doc) => {
        const lines = doc.lineCount;
        sendUpdate({ lines: lines })
    });

    vscode.languages.onDidChangeDiagnostics(() => {
        const diagnostics = vscode.languages.getDiagnostics()
        let errCount = 0;

        diagnostics.forEach(([_, diags]) => {
            errCount += diags.length;
        });

        sendUpdate({ errors: errCount });
    });
}

async function sendUpdate(data: any) {
    try {
        await fetch("http://localhost:8080/update", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },

            body: JSON.stringify(data),
        })
    } catch (err) {
        console.error(err);
    }
}

export function deactivate() {
    // TODO: finish later
}
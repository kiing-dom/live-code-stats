import * as vscode from 'vscode'

export function activate(context: vscode.ExtensionContext) {

    const keystrokeTimestamps: number[] = [];

    vscode.workspace.onDidChangeTextDocument((e) => {
        const now = Date.now();
        keystrokeTimestamps.push(now);

        const cutoff = now - 60000;
        while (keystrokeTimestamps.length > 0 && keystrokeTimestamps[0] < cutoff) {
            keystrokeTimestamps.shift();
        }

        sendUpdate({
            keystrokes: keystrokeTimestamps.length,
            file_name: e.document.fileName,
        });
    });

    vscode.workspace.onDidSaveTextDocument((doc) => {
        sendUpdate({ lines: doc.lineCount, file_name: doc.fileName });
    });

    vscode.languages.onDidChangeDiagnostics(() => {
        const activeEditor = vscode.window.activeTextEditor;
        if (!activeEditor) return;

        const activeFile = activeEditor.document.fileName;
        const diagnostics = vscode.languages.getDiagnostics(activeEditor.document.uri);
        const errCount = diagnostics.filter(d => d.severity === vscode.DiagnosticSeverity.Error).length;

        sendUpdate({
            errors: errCount,
            file_name: activeFile
        });
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
import MonacoEditor, { EditorDidMount } from 'react-monaco-editor';
import './CodeViewer.scss';

type Selection = {
  start: number;
  end: number;
};

type Props = {
  language?: string;
  theme?: string;
  value?: string;
  selections?: Selection[];
};

export default function CodeViewer(props: Props) {
  const editorDidMount: EditorDidMount = (editor, monaco) => {
    if (!props.selections || !props.selections.length) return;

    const options = {
      isWholeLine: true,
      inlineClassName: 'selected-line',
    };

    const selections = props.selections.map((i: Selection) => ({
      range: new monaco.Range(i.start, 1, i.end, 1),
      options,
    }));

    editor.deltaDecorations([], selections);

    const firstSelection = props.selections[0];
    editor.revealPositionInCenter({
      lineNumber: firstSelection.start,
      column: 1,
    });
  };

  return (
    <div className="code-viewer">
      <MonacoEditor
        language={props.language || 'json'}
        theme={props.theme || 'vs-dark'}
        value={props.value}
        options={{
          readOnly: true,
          selectOnLineNumbers: true,
          minimap: { enabled: false },
        }}
        editorDidMount={editorDidMount}
      />
    </div>
  );
}

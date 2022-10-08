import { HTMLAttributes } from 'react';
import DOMPurify from 'dompurify';
import { marked } from 'marked';
import classNames from '../../utils/className';
import './MarkdownViewer.scss';

type Props = HTMLAttributes<HTMLElement> & {
  text?: string;
};

export default function MarkdownViewer(props: Props) {
  const {
    text, className, ...other
  } = props;

  const fullClassName = classNames('markdown-viewer', className);

  if (!text) return (<div className={fullClassName} />);

  const dangerousHtml = marked(text, { gfm: true, breaks: true });
  const secureHtml = DOMPurify.sanitize(dangerousHtml);
  const html = { __html: secureHtml };

  const newProps = {
    ...other,
    className: fullClassName,
    dangerouslySetInnerHTML: html,
  };

  return <div {...newProps} />;
}

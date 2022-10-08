import { useState } from 'react';
import Red from './images/diff-deleted.png';
import { DiffData } from '../../../../../query/compare';

type Props = {
  data: DiffData.DiffDeletedItem;
};

export default function DiffDeletedItem(props: Props) {
  const [show, setShow] = useState(true);
  return (
    <div
      className="compare-row compare-row-deleted"
      onClick={(e) => {
        e.stopPropagation();
        setShow(!show);
      }}
    >
      <div className="row-item row-icon">
        {' '}
        <img className="icon" src={Red} alt="React Logo" />
      </div>
      <div className="row-item row-text">Deleted: </div>
      <div className="row-item row-code">
        {props.data?.method}
        {' '}
        {props.data?.path}
      </div>
      {props.data.breaking && <div className="row-breaking">Breaking</div>}
      {/* <div className="detail">
                  <MarkdownViewer text={i.message} />
              </div> */}
    </div>
  );
}

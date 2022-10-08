import { useState } from 'react';
import Green from './images/diff-added.png';
import { DiffData } from '../../../../../query/compare';

type Props = {
  data: DiffData.DiffAddedItem;
};

export default function DiffAddedItem(props: Props) {
  const [show, setShow] = useState(true);
  return (
    <div
      className="compare-row compare-row-added"
      onClick={(e) => {
        e.stopPropagation();
        setShow(!show);
      }}
    >
      <div className="row-item row-icon">
        {' '}
        <img className="icon" src={Green} alt="React Logo" />
      </div>
      <div className="row-item row-text">Added: </div>
      <div className="row-item row-code">
        {props.data?.method}
        {' '}
        {props.data?.path}
      </div>
      {props.data.breaking && <div className="row-breaking">Breaking</div>}
    </div>
  );
}

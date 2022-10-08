import DiffModifiedItem from './DiffModifiedItem/DiffModifiedItem';
import DiffAddedItem from './DiffAddItem/DiffAddItem';
import DiffDeletedItem from './DiffDeletedItem/DiffDeletedItem';
import { DiffData } from '../../../../query/compare';
import './DiffList.scss';

type Props = {
  data: DiffData.JsonDiffResult;
};

export default function DiffList(props: Props) {
  if (!props.data) {
    return (
      <div className="result-table">
        Select a service and two specs to compare
      </div>
    );
  }

  const addedRows = (props.data?.added || [])
    .map((addedItem, index) => {
      const key = `added-${index}`;
      return (
        <DiffAddedItem data={addedItem} key={key} />
      );
    });

  const modifiedRows = (props.data?.modified || [])
    .map((modifiedItem, index) => {
      const key = `modified-${index}`;
      return (
        <DiffModifiedItem data={modifiedItem} key={key} />
      );
    });

  const deletedRows = (props.data?.deleted || [])
    .map((deletedItem, index) => {
      const key = `deleted-${index}`;
      return (
        <DiffDeletedItem data={deletedItem} key={key} />
      );
    });

  const rows = [
    ...addedRows,
    ...modifiedRows,
    ...deletedRows,
  ];

  return (
    <div className="result-table">
      {rows}
    </div>
  );
}

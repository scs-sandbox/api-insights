import { ChangeEvent } from 'react';
import './AnalyzerFilter.scss';

type AnalyserItem = {
  status: string;
  title: string;
  value: string;
};

type AnalyserFilterItem = AnalyserItem & {
  selected: boolean;
};

type AnalyserFilterData = {
  [index: string]: AnalyserFilterItem;
};

type Props = {
  allItemEnabled?: boolean;
  analyzerList: AnalyserItem[];
  filterData: AnalyserFilterData;
  onChange: (data: AnalyserFilterData) => void;
};

export function filterIsSelected(
  filter: string,
  filterData?: AnalyserFilterData,
) {
  if (!filterData) return true;

  return filterData[filter] ? filterData[filter].selected : false;
}

function buildAllAnalyzerFilters(
  filterList: AnalyserItem[],
  allItemEnabled?: boolean,
) {
  if (!allItemEnabled) {
    return filterList || [];
  }

  return [
    {
      title: 'All',
      value: '*',
      status: '',
    },
    ...(filterList || []),
  ];
}

function buildNewFilterData(
  filterList: AnalyserItem[],
  oldFilterData: AnalyserFilterData,
) {
  if (!filterList) return oldFilterData;

  if (oldFilterData) return JSON.parse(JSON.stringify(oldFilterData));

  const allFilters = buildAllAnalyzerFilters(filterList);

  return allFilters.reduce(
    (pre: AnalyserFilterData, cur: AnalyserItem) => ({
      ...pre,
      [cur.value]: {
        ...cur,
        selected: true,
      },
    }),
    {},
  );
}

export function AnalyzerFilter(props: Props) {
  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    const newData = buildNewFilterData(props.analyzerList, props.filterData);
    const itemAll = newData['*'];
    const itemTarget = newData[e.currentTarget.dataset.value];

    if (itemAll === itemTarget) {
      const newDataValues = Object.values(newData) as AnalyserFilterItem[];
      for (let i = 0; i < newDataValues.length; i += 1) {
        const filterItem = newDataValues[i];
        filterItem.selected = e.currentTarget.checked;
      }
    } else {
      itemTarget.selected = e.currentTarget.checked;
      const noUnselectedItem = !Object.values(newData).find(
        (i: AnalyserFilterItem) => i.value !== '*' && !i.selected,
      );
      if (itemAll) {
        itemAll.selected = noUnselectedItem;
      }
    }

    if (props.onChange) {
      props.onChange(newData);
    }
  };

  if (!props.analyzerList) {
    return <div className="analyzer-filter" />;
  }

  const filterListItems = buildAllAnalyzerFilters(
    props.analyzerList,
    props.allItemEnabled,
  ).map((i) => {
    const checked = filterIsSelected(i.value, props.filterData);
    return (
      <li key={i.value} className="filter-item">
        <label className="item-label">
          <input
            className="input-part"
            type="checkbox"
            checked={checked}
            data-value={i.value}
            onChange={onChange}
          />
          <span
            className={`text-part ${
              i.status !== 'Analyzed' ? '' : 'analyze-failed'
            }`}
          >
            {i.status === 'Analyzed' ? i.title : `${i.title} (Failed)`}
          </span>
        </label>
      </li>
    );
  });

  return (
    <div className="analyzer-filter">
      <ul className="filter-list">{filterListItems}</ul>
    </div>
  );
}

import { MouseEvent, ReactElement } from 'react';
import classNames from '../../utils/className';
import './Tabs.scss';

type Props = {
  selectedTabIndex: number;
  headers: ReactElement[];
  children?: ReactElement[];
  onChangeIndex?: (index: number) => void;
};

export default function Tabs(props: Props) {
  const onClickHeader = (e: MouseEvent<HTMLElement>) => {
    const index = Number.parseInt(e.currentTarget.dataset.index, 10);

    if (props.onChangeIndex) {
      props.onChangeIndex(index);
    }
  };

  function renderHeaders() {
    const list = (props.headers || []).map((item, index) => {
      const className = classNames(
        'tab-header',
        index === props.selectedTabIndex ? 'active' : '',
      );

      const key = `key-${index}`;

      return (
        <div
          key={key}
          className={className}
          data-index={index}
          onClick={onClickHeader}
        >
          {item}
        </div>
      );
    });

    return <div className="tab-headers">{list}</div>;
  }

  const headers = renderHeaders();

  return (
    <div className="tabs">
      {headers}
      <div className="tab-bodys">{props.children}</div>
    </div>
  );
}

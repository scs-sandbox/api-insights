import { MouseEvent } from 'react';
import { ChartDataItem } from './TrendBar/Bar/Bar';
import TrendBar from './TrendBar/TrendBar';
import './TrendBarChart.scss';

type Props = {
  data: ChartDataItem[];
  minShareCount?: number;
  onEnterBar?: (e: MouseEvent<HTMLElement>) => void;
  onLeaveBar?: (e: MouseEvent<HTMLElement>) => void;
};

const SCALE = 20;
const SCALES = [100, 80, 60, 40, 20, 0];

export default function TrendBarChart(props: Props) {
  const getColumnsData = () => {
    const minShareCount = props.minShareCount || 0;
    const missingCount = minShareCount - props.data.length;
    if (missingCount <= 0) return props.data;

    return [...Array(missingCount).fill(null), ...props.data];
  };

  const renderAxisX = () => {
    const scales = SCALES.map((i: number, index: number) => {
      const style = {
        top: `${index * SCALE}%`,
      };

      const key = `key-${index}`;

      return (
        <div key={key} className="scale-value" style={style}>
          {i}
        </div>
      );
    });

    return (
      <div className="axis-x">
        <div className="scale-list">{scales}</div>
        <div className="axis-line" />
      </div>
    );
  };

  const renderBars = () => {
    if (!props.data) return null;
    if (!props.data.length) return null;

    const data = getColumnsData();

    const bars = data.map((i: ChartDataItem, index: number) => {
      const validItem = i && typeof i.score === 'number';
      const key = `key-${index}`;

      return (
        <div key={key} className="bar-col">
          {validItem && (
            <TrendBar
              {...i}
              onMouseEnter={props.onEnterBar}
              onMouseLeave={props.onLeaveBar}
            />
          )}
        </div>
      );
    });

    return <div className="bars">{bars}</div>;
  };

  const renderAxisXLines = () => {
    const lines = SCALES.map((i: number, index: number) => {
      const style = {
        top: `${index * SCALE}%`,
      };

      return <div key={i} style={style} className="axis-x-line block-item" />;
    });

    return <div className="axis-x-lines">{lines}</div>;
  };

  const renderChartArea = () => {
    const axisXLines = renderAxisXLines();
    const bars = renderBars();

    return (
      <div className="chart-area">
        {axisXLines}
        {bars}
      </div>
    );
  };

  const axisX = renderAxisX();
  const chartArea = renderChartArea();

  return (
    <div className="panel-block trend-bar-chart">
      <div className="chart-panel-box">
        <div className="chart-panel">
          {axisX}
          {chartArea}
        </div>
      </div>
    </div>
  );
}

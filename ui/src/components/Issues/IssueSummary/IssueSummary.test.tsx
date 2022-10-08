import { render } from '@testing-library/react';
import { ComplianceData } from '../../../query/compliance';
import IssueSummary from './IssueSummary';

describe('<IssueSummary />', () => {
  test('All issues', () => {
    const data: ComplianceData.ComplianceIssues[] = [{
      count: 1,
      error: {
        count: 1,
      },
      warning: {
        count: 10,
      },
      info: {
        count: 100,
      },
      hint: {
        count: 1000,
      },
    },
    {
      count: 1,
      error: {
        count: 1,
      },
      warning: {
        count: 10,
      },
      info: {
        count: 100,
      },
      hint: {
        count: 1000,
      },
    }];

    const { container } = render((
      <IssueSummary data={data} showLabel />
    ));

    const errorItem = container.querySelector('.issue-item-error');
    expect(errorItem).toBeInTheDocument();
    const errorItemCount = errorItem.querySelector('.issue-item-count');
    expect(errorItemCount).toBeInTheDocument();
    expect(errorItemCount).toHaveTextContent('2');
    expect(errorItem.querySelector('.issue-item-label')).toHaveTextContent('Error');

    const warningItem = container.querySelector('.issue-item-warning');
    expect(warningItem).toBeInTheDocument();
    const warningItemCount = warningItem.querySelector('.issue-item-count');
    expect(warningItemCount).toBeInTheDocument();
    expect(warningItemCount).toHaveTextContent('20');
    expect(warningItem.querySelector('.issue-item-label')).toHaveTextContent('Warning');

    const infoItem = container.querySelector('.issue-item-info');
    expect(infoItem).toBeInTheDocument();
    const infoItemCount = infoItem.querySelector('.issue-item-count');
    expect(infoItemCount).toBeInTheDocument();
    expect(infoItemCount).toHaveTextContent('200');
    expect(infoItem.querySelector('.issue-item-label')).toHaveTextContent('Info');

    const hintItem = container.querySelector('.issue-item-hint');
    expect(hintItem).toBeInTheDocument();
    const hintItemCount = hintItem.querySelector('.issue-item-count');
    expect(hintItemCount).toBeInTheDocument();
    expect(hintItemCount).toHaveTextContent('2000');
    expect(hintItem.querySelector('.issue-item-label')).toHaveTextContent('Hint');
  });

  test('Data is Null', () => {
    const data: ComplianceData.ComplianceIssues[] = null;

    const { container } = render((
      <IssueSummary data={data} showLabel />
    ));

    const errorItem = container.querySelector('.issue-item-error');
    expect(errorItem).toBeInTheDocument();
    const errorItemCount = errorItem.querySelector('.issue-item-count');
    expect(errorItemCount).toBeInTheDocument();
    expect(errorItemCount).toHaveTextContent('0');
    expect(errorItem.querySelector('.issue-item-label')).toHaveTextContent('Error');

    const warningItem = container.querySelector('.issue-item-warning');
    expect(warningItem).toBeInTheDocument();
    const warningItemCount = warningItem.querySelector('.issue-item-count');
    expect(warningItemCount).toBeInTheDocument();
    expect(warningItemCount).toHaveTextContent('0');
    expect(warningItem.querySelector('.issue-item-label')).toHaveTextContent('Warning');

    const infoItem = container.querySelector('.issue-item-info');
    expect(infoItem).toBeInTheDocument();
    const infoItemCount = infoItem.querySelector('.issue-item-count');
    expect(infoItemCount).toBeInTheDocument();
    expect(infoItemCount).toHaveTextContent('0');
    expect(infoItem.querySelector('.issue-item-label')).toHaveTextContent('Info');

    const hintItem = container.querySelector('.issue-item-hint');
    expect(hintItem).toBeInTheDocument();
    const hintItemCount = hintItem.querySelector('.issue-item-count');
    expect(hintItemCount).toBeInTheDocument();
    expect(hintItemCount).toHaveTextContent('0');
    expect(hintItem.querySelector('.issue-item-label')).toHaveTextContent('Hint');
  });

  test('Some issues', () => {
    const data: ComplianceData.ComplianceIssues[] = [{
      count: 1,
      error: {
        count: 1,
      },
    },
    {
      count: 1,
      error: {
        count: 1,
      },
      warning: {
        count: 10,
      },
    }];

    const { container } = render((
      <IssueSummary data={data} showLabel />
    ));

    const errorItem = container.querySelector('.issue-item-error');
    expect(errorItem).toBeInTheDocument();
    const errorItemCount = errorItem.querySelector('.issue-item-count');
    expect(errorItemCount).toBeInTheDocument();
    expect(errorItemCount).toHaveTextContent('2');
    expect(errorItem.querySelector('.issue-item-label')).toHaveTextContent('Error');

    const warningItem = container.querySelector('.issue-item-warning');
    expect(warningItem).toBeInTheDocument();
    const warningItemCount = warningItem.querySelector('.issue-item-count');
    expect(warningItemCount).toBeInTheDocument();
    expect(warningItemCount).toHaveTextContent('10');
    expect(warningItem.querySelector('.issue-item-label')).toHaveTextContent('Warning');

    const infoItem = container.querySelector('.issue-item-info');
    expect(infoItem).toBeInTheDocument();
    const infoItemCount = infoItem.querySelector('.issue-item-count');
    expect(infoItemCount).toBeInTheDocument();
    expect(infoItemCount).toHaveTextContent('0');
    expect(infoItem.querySelector('.issue-item-label')).toHaveTextContent('Info');

    const hintItem = container.querySelector('.issue-item-hint');
    expect(hintItem).toBeInTheDocument();
    const hintItemCount = hintItem.querySelector('.issue-item-count');
    expect(hintItemCount).toBeInTheDocument();
    expect(hintItemCount).toHaveTextContent('0');
    expect(hintItem.querySelector('.issue-item-label')).toHaveTextContent('Hint');
  });
});

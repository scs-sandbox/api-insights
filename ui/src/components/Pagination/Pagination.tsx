import Pages from 'react-js-pagination';
import './Pagination.scss';

type Props = {
  currentPage: number;
  pageSize: number;
  total: number;
  pageRangeDisplayed?: number;
  prevPageText?: string;
  nextPageText?: string;
  firstPageText?: string;
  lastPageText?: string;
  onChange?: (value: number) => void;
};

export default function Pagination(props: Props) {
  const onChange = (e: number) => {
    const index = e;
    if (props.onChange) {
      props.onChange(index);
    }
  };

  return (
    <Pages
      prevPageText={props.prevPageText || '<'}
      nextPageText={props.nextPageText || '>'}
      firstPageText={props.firstPageText || '<<'}
      lastPageText={props.lastPageText || '>>'}
      itemClassFirst="move-btn first-page"
      itemClassPrev="move-btn prev-page"
      itemClassNext="move-btn next-page"
      itemClassLast="move-btn last-page"
      itemClass="page-btn"
      linkClass="page-btn-link"
      activePage={props.currentPage}
      itemsCountPerPage={props.pageSize}
      totalItemsCount={props.total}
      pageRangeDisplayed={props.pageRangeDisplayed || 10}
      onChange={onChange}
    />
  );
}

import dayjs from "dayjs";

const formatDate = (date, format = 'YYYY-MM-DD') => {
  if (!date) {
    return '';
  }

  return dayjs(date).format(format)
};

export default function (Vue) {
  Vue.filter('formatDate', function (value, format = 'YYYY-MM-DD') {
    return formatDate(value, format);
  });
}
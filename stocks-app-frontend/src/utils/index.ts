import { type FunctionalComponent, type HTMLAttributes, type VNodeProps } from 'vue';
import { ArrowUpIcon, ArrowDownIcon, ArrowPathIcon } from '@heroicons/vue/24/solid';

export const formatDate = (dateString: string) => {
  try {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  } catch (e) {
    return 'Invalid date';
  }
};

export const actionIcons: { [key: string]: FunctionalComponent<HTMLAttributes & VNodeProps, {}, any, {}> } = {
  'target raised by': ArrowUpIcon,
  'upgraded by': ArrowUpIcon,
  'target lowered by': ArrowDownIcon,
  'downgraded by': ArrowDownIcon,
  'reiterated by': ArrowPathIcon,
  'initiated by': ArrowPathIcon,
};

export const actionColors: { [key: string]: string } = {
  'target raised by': 'text-green-600',
  'upgraded by': 'text-green-600',
  'target lowered by': 'text-red-600',
  'downgraded by': 'text-red-600',
  'reiterated by': 'text-blue-600',
  'initiated by': 'text-purple-600',
};

export const ratingColors: { [key: string]: string } = {
  'Buy': 'bg-green-100 text-green-800',
  'Outperform': 'bg-green-100 text-green-800',
  'Overweight': 'bg-green-100 text-green-800',
  'Equal Weight': 'bg-blue-100 text-blue-800',
  'Sector Perform': 'bg-blue-100 text-blue-800',
  'Hold': 'bg-yellow-100 text-yellow-800',
  'Neutral': 'bg-yellow-100 text-yellow-800',
  'Underperform': 'bg-red-100 text-red-800',
  'Sell': 'bg-red-100 text-red-800',
};

export const getChangePercentClass = (changePercent: number) => {
  return changePercent >= 0 ? 'text-green-600' : 'text-red-600';
};
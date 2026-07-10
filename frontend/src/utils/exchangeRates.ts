export const exchangeRates: Record<string, Record<string, number>> = {
  USD: {
    USD: 1,
    NGN: 1550,
    EUR: 0.92,
    GBP: 0.79,
  },

  NGN: {
    NGN: 1,
    USD: 1 / 1550,
  },

  EUR: {
    EUR: 1,
    USD: 1 / 0.92,
  },

  GBP: {
    GBP: 1,
    USD: 1 / 0.79,
  },
};
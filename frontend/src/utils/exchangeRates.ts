export const exchangeRates: Record<string, Record<string, number>> = {
  NGN: {
    NGN: 1,
    USD: 0.00065,
    EUR: 0.00055,
  },

  USD: {
    USD: 1,
    NGN: 1550,
    EUR: 0.85,
  },

  EUR: {
    EUR: 1,
    USD: 1.17,
    NGN: 1820,
  },
};
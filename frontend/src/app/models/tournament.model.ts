export interface Tournament {
  id: number;
  name: string;
  locationName: string;
  streetAddress: string;
  city: string;
  state: string;
  organizerId: number;
  startDate: Date;
  endDate: Date;
  boysVarsity: number;
  girlsVarsity: number;
  boysJv: number;
  girlsJv: number;
  boysMs: number;
  girlsMs: number;
  boysYouth: number;
  girlsYouth: number;
}

export const AgeDivisions = [
  "boysVarsity",
  "girlsVarsity",
  "boysJv",
  "girlsJv",
  "boysMs",
  "girlsMs",
  "boysYouth",
  "girlsYouth",
];

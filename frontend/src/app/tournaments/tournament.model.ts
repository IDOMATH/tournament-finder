export interface Tournament {
  id: number;
  name: string;
  locationName: string;
  streetAddress: string;
  city: string;
  state: string;
  organizerId: number;
  isFull: boolean;
  startDate: Date;
  endDate: Date;
  isBoysVarsity: boolean;
  isGirlsVarsity: boolean;
  isBoysJv: boolean;
  isGirlsJv: boolean;
  isBoysMs: boolean;
  isGirlsMs: boolean;
  isBoysYouth: boolean;
  isGirlsYouth: boolean;
}

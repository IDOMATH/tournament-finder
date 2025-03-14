export interface Tournament {
  id: number;
  name: string;
  locationName: string;
  locationAddress: string;
  organizerId: number;
  isFull: boolean;
  startDate: Date;
  endDate: Date;
  ageDivision: boolean[]; //This should only ever be 8 in length
}

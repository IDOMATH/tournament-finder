import { AbstractControl } from "@angular/forms";

export function endDateNotBeforeStartDate(
  startControl: string,
  endControl: string
) {
  return (control: AbstractControl) => {
    const startDateString = control.get(startControl)?.value;
    const endDateString = control.get(endControl)?.value;

    const startDate = new Date(startDateString);
    const endDate = new Date(endDateString);

    if (endDate < startDate) {
      return { endDateBeforeStartDate: true };
    }

    return null;
  };
}

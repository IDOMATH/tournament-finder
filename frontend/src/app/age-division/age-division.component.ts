import { Component, input } from "@angular/core";

@Component({
  selector: "app-age-divion",
  imports: [],
  templateUrl: "./age-division.component.html",
  styleUrl: "./age-division.component.css",
})
export class AgeDivisionComponent {
  ageDivision = input.required<string>();
  availability = input.required<number>();

  getAvailabilityText() {
    if (this.availability() === 0) {
      return "not offered.";
    }
    if (this.availability() === 1) {
      return "taking registrations.";
    }
    if (this.availability() === 2) {
      return "full.";
    }

    //TODO: this should probably log something somewhere
    return "has an error.  Let us know how you found this.";
  }
}

import { Component, input } from "@angular/core";

@Component({
  selector: "app-age-divion",
  imports: [],
  templateUrl: "./age-division.component.html",
  styleUrl: "./age-division.component.css",
})
export class AgeDivisionComponent {
  ageDivision = input.required<number>();
}

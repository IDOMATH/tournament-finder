import { DatePipe } from "@angular/common";
import { Component, input } from "@angular/core";

@Component({
  selector: "app-date-display",
  imports: [DatePipe],
  templateUrl: "./date-display.component.html",
  styleUrl: "./date-display.component.css",
})
export class DateDisplayComponent {
  startDate = input.required<Date>();
  endDate = input.required<Date>();
}

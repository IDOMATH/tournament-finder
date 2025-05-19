import { Component, input } from "@angular/core";
import { DateDisplayComponent } from "../../date-display/date-display.component";
import { RouterLink } from "@angular/router";

@Component({
  selector: "app-tournament-card",
  imports: [DateDisplayComponent, RouterLink],
  templateUrl: "./tournament-card.component.html",
  styleUrl: "./tournament-card.component.css",
})
export class TournamentCardComponent {
  name = input.required<string>();
  startDate = input.required<Date>();
  endDate = input.required<Date>();
  id = input.required<number>();
}

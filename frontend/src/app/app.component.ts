import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { NewTournamentFormComponent } from "./new-tournament-form/new-tournament-form.component";

@Component({
  selector: "app-root",
  imports: [RouterOutlet, NewTournamentFormComponent],
  templateUrl: "./app.component.html",
  styleUrl: "./app.component.css",
})
export class AppComponent {
  title = "tournament-finder";
}

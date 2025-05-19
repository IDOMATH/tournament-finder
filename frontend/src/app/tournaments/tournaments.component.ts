import { Component, inject } from "@angular/core";
import { TournamentService } from "../services/tournament-service";
import { TournamentCardComponent } from "./tournament-card/tournament-card.component";

@Component({
  selector: "app-tournaments",
  imports: [TournamentCardComponent],
  templateUrl: "./tournaments.component.html",
  styleUrl: "./tournaments.component.css",
})
export class TournamentsComponent {
  private tournamentService = inject(TournamentService);

  tournaments = this.tournamentService.getTournaments();
}

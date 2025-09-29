import { Component, DestroyRef, inject } from "@angular/core";
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
  private destroyRef = inject(DestroyRef);

  tournaments = this.tournamentService.loadedTournaments;

  ngOnInit(): void {
    const subscription = this.tournamentService
      .setFetchedTournaments()
      .subscribe();

    // clean up subscription with destroyRef
  }
}

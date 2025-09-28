import { inject, Injectable, signal } from "@angular/core";
import { Tournament } from "../models/tournament.model";
import { HttpClient } from "@angular/common/http";
import { map, tap } from "rxjs";
import { outputToObservable } from "@angular/core/rxjs-interop";

@Injectable({ providedIn: "root" })
export class TournamentService {
  private httpClient = inject(HttpClient);
  private tournament = signal<Tournament | undefined>(undefined);
  private tournaments = signal<Tournament[]>([]);

  loadedTournaments = this.tournaments.asReadonly();

  setFetchedTournaments() {
    return this.getTournaments().pipe(
      tap({
        next: (tournaments) => this.tournaments.set(tournaments),
      })
    );
  }

  getTournamentById(id: number) {
    return this.httpClient
      .get<{ tournament: Tournament }>(
        "http://localhost:8080/tournaments/" + id
      )
      .pipe(map((resData) => this.tournament.set(resData.tournament)));
  }

  getTournaments() {
    return this.httpClient
      .get<{ tournaments: Tournament[] }>("http://localhost:8080/tournaments/")
      .pipe(map((resData) => resData.tournaments));
  }

  postTournament(tournament: Tournament) {
    JSON.stringify(tournament);
    // call to post tournament
    this.httpClient.post("http://localhost:8080/tournaments/", tournament);
  }
  putTournament(tournament: Tournament) {
    JSON.stringify(tournament);
    this.httpClient.put("http://localhost:8080/tournaments/", tournament);
  }
  deleteTournament(id: number) {
    this.httpClient.delete("http://localhost:8080/tournaments/" + id);
  }
}

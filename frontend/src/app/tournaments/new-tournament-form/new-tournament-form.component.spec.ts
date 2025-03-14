import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewTournamentFormComponent } from './new-tournament-form.component';

describe('NewTournamentFormComponent', () => {
  let component: NewTournamentFormComponent;
  let fixture: ComponentFixture<NewTournamentFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NewTournamentFormComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NewTournamentFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

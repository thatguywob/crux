package agent

func (m *Manager) bobPreflightUnlocked(agent *AgentProcess) {
	if agent == nil {
		return
	}
	backend := agent.Config.Backend
	if agent.BackendOverride != "" {
		backend = agent.BackendOverride
	}
	if backend != bobBackend || m.bobAPIKey() == "" {
		return
	}
	// Bob shared-home repair and readability probes touch /data/home and run
	// subprocesses. Keep them in unlocked prep phases so a slow PVC or UID
	// probe cannot block AllStatuses()/heartbeat while boot launches stagger
	// across the agent fleet.
	m.ensureBobAuthSettings(agent.Name, bobSharedHome)
	m.verifyBobKeyReadable(agent.Name, m.bobKeyFilePath(), agent.UID)
	_ = m.verifyBobStateDirsWritable(agent.Name, bobSharedHome, m.workDir+"/"+agent.Name, agent.UID)
}
